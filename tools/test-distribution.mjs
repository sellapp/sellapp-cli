import assert from 'node:assert/strict';
import { createHash } from 'node:crypto';
import { chmodSync, existsSync, mkdirSync, mkdtempSync, readFileSync, realpathSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join, resolve } from 'node:path';
import { spawnSync } from 'node:child_process';

const version = '0.1.1';
assert.equal(process.argv[2], '--artifacts', 'Usage: node tools/test-distribution.mjs --artifacts DIRECTORY');
assert.equal(process.argv.length, 4, 'Specify exactly one artifact directory');
const artifacts = resolve(process.argv[3]);
const root = realpathSync(mkdtempSync(join(tmpdir(), 'sellapp-distribution-')));
const home = join(root, 'home');
mkdirSync(home);
const env = { ...process.env, HOME: home, USERPROFILE: home, XDG_CONFIG_HOME: join(home, '.config'), APPDATA: join(home, 'AppData'), SELLAPP_AUDIT_DISABLED: '1', npm_config_cache: join(root, 'npm-cache'), npm_config_userconfig: join(root, 'npmrc'), npm_config_offline: 'true', npm_config_ignore_scripts: 'true', npm_config_audit: 'false', npm_config_fund: 'false' };
for (const key of Object.keys(env)) { if (key.startsWith('SELLAPP_') && key !== 'SELLAPP_AUDIT_DISABLED') delete env[key]; }
const skippedChecks = ['Native execution of foreign target binaries', 'Live OAuth login, refresh and disconnect', 'Registry publication and hosted installer'];
let passedTests = 0;
const run = (command, args, options = {}) => {
  const result = spawnSync(command, args, { cwd: root, env, encoding: 'utf8', timeout: 120000, ...options });
  assert.equal(result.status, 0, command + ' failed:\n' + (result.error?.message || '') + result.stdout + result.stderr);
  return result;
};
const test = (name, action) => { action(); passedTests++; process.stderr.write('PASS ' + name + '\n'); };
const sha = file => createHash('sha256').update(readFileSync(file)).digest('hex');
const executable = (file, content) => { writeFileSync(file, content); chmodSync(file, 0o755); };
let installerVerified = false;
try {
  const main = join(artifacts, 'sellapp-cli-' + version + '.tgz');
  const platform = join(artifacts, 'sellapp-cli-' + process.platform + '-' + process.arch + '-' + version + '.tgz');
  test('all eight npm tarballs exist', () => {
    for (const target of ['linux-x64', 'linux-arm64', 'darwin-x64', 'darwin-arm64', 'win32-ia32', 'win32-x64', 'win32-arm64']) assert.ok(existsSync(join(artifacts, 'sellapp-cli-' + target + '-' + version + '.tgz')));
    assert.ok(existsSync(main));
  });
  const npmCLI = [process.env.npm_execpath, join(dirname(process.execPath), 'node_modules/npm/bin/npm-cli.js'), join(dirname(process.execPath), '../lib/node_modules/npm/bin/npm-cli.js')].find(file => file && existsSync(file));
  const npmRun = args => npmCLI ? run(process.execPath, [npmCLI, ...args]) : run('npm', args);
  const installArgs = ['--offline', '--ignore-scripts', '--no-audit', '--no-fund', '--include=optional'];
  for (const global of [false, true]) {
    const prefix = join(root, global ? 'global prefix' : 'local project');
    mkdirSync(prefix);
    test(global ? 'global npm installation from local tarballs' : 'local npm installation from local tarballs', () => {
      npmRun(['install', ...(global ? ['--global'] : []), '--prefix', prefix, ...installArgs, main, platform]);
      const launcher = process.platform === 'win32'
        ? join(prefix, ...(global ? [] : ['node_modules', '.bin']), 'sellapp.cmd')
        : join(prefix, ...(global ? ['bin'] : ['node_modules', '.bin']), 'sellapp');
      assert.ok(existsSync(launcher));
      const packageRoot = join(prefix, ...(global && process.platform !== 'win32' ? ['lib'] : []), 'node_modules', '@sell.app', 'cli');
      const invoke = args => process.platform === 'win32' ? run(process.execPath, [join(packageRoot, 'bin/sellapp.cjs'), ...args]) : run(launcher, args);
      const result = invoke(['--version']);
      assert.equal(result.stdout.trim(), 'sellapp version ' + version);
      assert.equal(result.stderr, '');
      const metadata = JSON.parse(readFileSync(join(packageRoot, 'package.json'), 'utf8'));
      assert.equal(metadata.version, version);
      assert.equal(metadata.scripts, undefined);
      for (const file of ['skills/sellapp/SKILL.md', 'docs/usage.md', 'docs/contributing.md', 'LICENSE.txt', 'README.md']) assert.ok(readFileSync(join(packageRoot, file)).length > 0);
      const resultHelp = invoke(['--help']);
      assert.ok(resultHelp.stdout.includes('sellapp'));
    });
  }
  if (process.platform === 'linux' || process.platform === 'darwin') {
    const tools = join(root, 'tools');
    mkdirSync(tools);
    const arch = process.arch === 'x64' ? 'amd64' : process.arch;
    const asset = 'sellapp_' + version + '_' + process.platform + '_' + arch + '.tar.gz';
    const archive = join(artifacts, asset);
    const sums = join(root, 'SHA256SUMS');
    const log = join(root, 'downloads');
    writeFileSync(sums, readFileSync(join(artifacts, 'SHA256SUMS')));
    executable(join(tools, 'curl'), '#!/bin/sh\nset -eu\nprintf "%s\\n" "$*" >> "$TEST_DOWNLOAD_LOG"\noutput=""\nwhile [ "$#" -gt 0 ]; do case "$1" in --output) output=$2; shift 2 ;; *) url=$1; shift ;; esac; done\ncase "$url" in */latest) printf "%s" "https://github.com/sellapp/sellapp-cli/releases/tag/v0.1.1" ;; */SHA256SUMS) cp "$TEST_CHECKSUMS" "$output" ;; */"$TEST_ASSET") cp "$TEST_ARCHIVE" "$output" ;; *) exit 64 ;; esac\n');
    const destination = join(home, "bin with space's");
    const shellProfile = process.platform === 'darwin' ? '.bash_profile' : '.bashrc';
    const installEnv = { ...env, PATH: tools + ':' + process.env.PATH, SHELL: '/bin/bash', SELLAPP_INSTALL_DIR: destination, SELLAPP_INSTALL_VERSION: version, TEST_DOWNLOAD_LOG: log, TEST_CHECKSUMS: sums, TEST_ARCHIVE: archive, TEST_ASSET: asset };
    const install = overrides => run('sh', [join(artifacts, 'install.sh')], { env: { ...installEnv, ...overrides } });
    test('installer verifies and executes the same native release binary', () => {
      install({});
      const expected = spawnSync('tar', ['-xOzf', archive, 'sellapp'], { maxBuffer: 128 * 1024 * 1024 });
      assert.equal(expected.status, 0);
      assert.equal(sha(join(destination, 'sellapp')), createHash('sha256').update(expected.stdout).digest('hex'));
      assert.equal(run(join(destination, 'sellapp'), ['--version']).stdout.trim(), 'sellapp version ' + version);
      assert.ok(readFileSync(join(home, shellProfile), 'utf8').includes('SellApp CLI'));
    });
    test('pinned/latest and repeated installations preserve one escaped PATH block', () => {
      install({ SELLAPP_INSTALL_VERSION: 'latest' });
      install({});
      assert.equal((readFileSync(log, 'utf8').match(/\/releases\/latest/g) || []).length, 1);
      assert.equal((readFileSync(join(home, shellProfile), 'utf8').match(/# >>> SellApp CLI >>>/g) || []).length, 1);
      const discovered = run('sh', ['-c', '. "$HOME/' + shellProfile + '"; command -v sellapp'], { env: installEnv });
      assert.equal(discovered.stdout.trim(), join(destination, 'sellapp'));
    });
    test('checksum rejection preserves the current installation', () => {
      const before = sha(join(destination, 'sellapp'));
      writeFileSync(sums, '0'.repeat(64) + '  ' + asset + '\n');
      const result = spawnSync('sh', [join(artifacts, 'install.sh')], { env: installEnv, encoding: 'utf8' });
      assert.notEqual(result.status, 0);
      assert.ok(result.stderr.includes('checksum did not match'));
      assert.equal(sha(join(destination, 'sellapp')), before);
    });
    test('interrupted replacement preserves the current installation', () => {
      writeFileSync(sums, readFileSync(join(artifacts, 'SHA256SUMS')));
      const before = sha(join(destination, 'sellapp'));
      executable(join(tools, 'mv'), '#!/bin/sh\nkill -TERM "$PPID"\nexit 1\n');
      const result = spawnSync('sh', [join(artifacts, 'install.sh')], { env: installEnv, encoding: 'utf8' });
      assert.notEqual(result.status, 0);
      assert.equal(sha(join(destination, 'sellapp')), before);
      rmSync(join(tools, 'mv'));
    });
    installerVerified = true;
  } else skippedChecks.push('POSIX installer is not supported on this host');
  process.stdout.write(JSON.stringify({ npmPackages: 8, npmInstallValidated: true, installerVerified, passedTests, nativePlatform: process.platform, nativeArchitecture: process.arch, skippedChecks }) + '\n');
} finally { rmSync(root, { recursive: true, force: true }); }
