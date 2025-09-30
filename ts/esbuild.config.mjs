import * as esbuild from 'esbuild';

/**
 * @param {esbuild.BuildOptions} config
 */
const baseConfig = (config) => ({
  ...config,
  entryPoints: [
    `src/index.ts`,
    `src/index.web.ts`,
    'src/generated/protos/index.*',
  ],
  bundle: true,
  sourcemap: true,
  packages: "external",
  platform: "neutral",
  external: [
    "node:*",
  ],
  outExtension: config.format === 'cjs' ? { '.js': '.cjs' } : undefined,
  minify: false,
  target: [`es2020`],
  splitting: config.format === 'esm',
  outdir: `dist/${config.format}`,
});


await Promise.all([
  esbuild.build(baseConfig({ format: 'esm' })),
  esbuild.build(baseConfig({ format: 'cjs' })),
]);
console.log('Building JS SDK finished');
