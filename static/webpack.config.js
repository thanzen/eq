module.exports = {
  context: __dirname,
  entry: './admin.js',
  output: {
    filename: 'dist/admin.js'
  },
  resolve: {
    extensions: ['', '.webpack.js', '.web.js', '.js', '.ts']
  },
  // devtool: "#inline-source-map", // Jup, sourcemaps
  module: {
    loaders: [{
      test: /\.js$/,
      loader: 'jsx-loader'
    }]
  }
};
