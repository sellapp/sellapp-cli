#!/usr/bin/env node
'use strict';
const { spawn } = require('node:child_process');
const path = require('node:path');
function fail(message) { process.stderr.write('sellapp: ' + message + '\n'); process.exitCode = 1; }
if (Number(process.versions.node.split('.')[0]) < 22) {
  fail('The npm launcher requires Node.js 22 or newer. Upgrade Node.js or use the standalone installer.');
} else {
  const target = process.platform + '-' + process.arch;
  const supported = ["linux-x64","linux-arm64","darwin-x64","darwin-arm64","win32-ia32","win32-x64","win32-arm64"];
  if (!supported.includes(target)) {
    fail('Unsupported platform ' + target + '. Supported platforms: ' + supported.join(', ') + '.');
  } else {
    const name = '@sell.app/cli-' + target;
    let executable;
    try {
      const manifestPath = require.resolve(name + '/package.json');
      const manifest = require(manifestPath);
      if (manifest.version !== '0.1.1') throw new Error('platform package version does not match');
      executable = path.join(path.dirname(manifestPath), 'bin', process.platform === 'win32' ? 'sellapp.exe' : 'sellapp');
      require('node:fs').accessSync(executable, require('node:fs').constants.X_OK);
    } catch {
      fail('Missing or incompatible ' + name + '@0.1.1. Reinstall with npm install -g @sell.app/cli@0.1.1 --include=optional (or omit -g for a local install). Optional dependencies must be enabled; the launcher does not download binaries.');
    }
    if (executable) {
      const child = spawn(executable, process.argv.slice(2), { stdio: 'inherit', windowsHide: true });
      const handlers = new Map();
      for (const signal of ['SIGINT', 'SIGTERM', 'SIGHUP']) {
        const handler = () => { if (child.exitCode === null && child.signalCode === null) child.kill(signal); };
        handlers.set(signal, handler);
        process.on(signal, handler);
      }
      const cleanup = () => { for (const [signal, handler] of handlers) process.removeListener(signal, handler); };
      child.on('error', (error) => { cleanup(); fail('Could not start ' + executable + ': ' + error.message + '. Reinstall @sell.app/cli.'); });
      child.on('exit', (code, signal) => {
        cleanup();
        if (signal) { try { process.kill(process.pid, signal); } catch { process.exitCode = 1; } }
        else process.exitCode = code === null ? 1 : code;
      });
    }
  }
}
