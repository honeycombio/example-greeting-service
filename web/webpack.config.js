const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const webpack = require("webpack");

module.exports = {
  entry: {
    index: './src/index.js',
    'tracing-http': './src/tracing-http.js',
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
      new webpack.DefinePlugin({
          FRONTEND_ENDPOINT: JSON.stringify(process.env.FRONTEND_ENDPOINT || 'http://localhost:7777'),
      }),
  ],
  resolve: {
    fallback: {
      path: require.resolve('path-browserify'),
    },
  },
};
