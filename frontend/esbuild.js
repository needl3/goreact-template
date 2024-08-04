const postCssPlugin = require('esbuild-style-plugin')
const esbuild = require('esbuild')

const config = {
  entryPoints: ['src/index.tsx', 'src/style.css'],
  outdir: 'public/assets',
  bundle: true,
  minify: false,
  plugins: [
    postCssPlugin({
      postcss: {
        plugins: [require('tailwindcss'), require('autoprefixer')],
      },
    }),
  ],
}

esbuild
  .build(config)
  .catch(() => {
    console.error(`Build error: ${error}`)
    process.exit(1)
  }).then(() => {
    const run = async () => {
      const ctx = await esbuild.context(config);
      await ctx.watch();
    };

    run();
    console.log("Build success. Watching for changes...")
  })
