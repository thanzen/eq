var path = require("path");
var CommonsChunkPlugin = require("./node_modules/webpack/lib/optimize/CommonsChunkPlugin");
var ExtractTextPlugin = require("extract-text-webpack-plugin");
module.exports = {
  context: __dirname,
  entry: {
    admin:'./admin.js',
  },
  output: {
       path: path.join(__dirname, "dist"),
       filename: "[name].js"
   },
  resolve: {
    extensions: ['', '.webpack.js', '.web.js', '.js', '.ts']
  },
  devtool: "#inline-source-map", // Jup, sourcemaps
  module: {
    loaders: [{
      test: /\.js$/,
      loader: 'jsx-loader'
    },
    { test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "url-loader?limit=10000&minetype=application/font-woff" },
      { test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/, loader: "file-loader" },
          {
                   test: /\.css$/,
                   loader: ExtractTextPlugin.extract("style-loader", "css-loader")
               },
    ]
  },
    plugins: [
        // new CommonsChunkPlugin("commons", "commons.js", ["admin"]),
        new ExtractTextPlugin("[name].css"),
    ]
};
