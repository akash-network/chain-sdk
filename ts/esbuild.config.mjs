import * as esbuild from 'esbuild';
import packageDetails from './package.json' with { type: 'json' };

/**
 * @param {"server"|"web"} type
 * @param {esbuild.BuildOptions} config
 */
const baseConfig = (type, config) => ({
  ...config,
  entryPoints: [
    `src/sdk/chain/index.${type}.ts`,
    `src/sdk/provider/index.${type}.ts`,
    'src/sdl/index.ts',
    'src/generated/protos/index.*'
  ],
  bundle: true,
  sourcemap: true,
  external: Object.keys(packageDetails.dependencies),
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

await Promise.all([
  esbuild.build(nodeJsConfig('esm')),
  esbuild.build(nodeJsConfig('cjs')),
]);
console.log('Building Nodejs SDK finished');
