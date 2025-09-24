import * as esbuild from 'esbuild';
import packageDetails from './package.json' with { type: 'json' };

/**
 * @param {"server"|"web"} type
 * @param {esbuild.BuildOptions} config
 */
const baseConfig = (type, config) => ({
  ...config,
  entryPoints: [
    `src/sdk/chain/${type}/index.ts`,
    `src/sdk/provider/${type}/index.ts`,
    'src/sdl/index.ts',
    'src/generated/protos/index.*'
  ],
  bundle: true,
  sourcemap: true,
  packages: "external",
  external: [
    "node:*",
  ],
});

/**
 * @type {esbuild.BuildOptions}
 * @param {esbuild.BuildOptions['format']} format
 */
const nodeJsConfig = (format) => baseConfig('server', {
  minify: false,
  target: [`node${packageDetails.engines.node}`],
  format,
  splitting: format === 'esm',
  platform: 'node',
  outdir: `dist/nodejs/${format}`,
});

const webConfig = (format) => baseConfig('web', {
  minify: false,
  target: ['es2020'],
  format,
  splitting: format === 'esm',
  platform: 'browser',
  outdir: `dist/web/${format}`,
});

await Promise.all([
  esbuild.build(nodeJsConfig('esm')),
  esbuild.build(nodeJsConfig('cjs')),
  esbuild.build(webConfig('esm')),
  esbuild.build(webConfig('cjs')),
]);
console.log('Building JS SDK finished');
