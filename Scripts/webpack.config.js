module.exports = {
  context: __dirname,
  entry: './app.js',
  output: {
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['', '.webpack.js', '.web.js', '.js', '.ts']
  },
  devtool: "#inline-source-map", // Jup, sourcemaps
  module: {
    loaders: [{
      test: /\.js$/,
      loader: 'jsx-loader'
    }]
  }
};
