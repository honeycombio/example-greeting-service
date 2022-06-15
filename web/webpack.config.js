const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: {
    "tracing-http": './src/tracing-http.js',
  },
  devtool: 'inline-source-map',
  devServer: {
    static: './bld',
  },
  mode: 'development',
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'dist'),
    clean: true,
  },
  optimization: {
    runtimeChunk: 'single',
  },
  plugins: [
    new HtmlWebpackPlugin({
      title: 'Development',
    }),
  ],
  resolve: {
    fallback: {
      path: require.resolve("path-browserify")
    }
  }
};