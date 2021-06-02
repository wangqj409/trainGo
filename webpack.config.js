let fs = require('fs');
let path = require('path');
let HtmlwebpackPlugin = require('html-webpack-plugin')
let glob = require('glob')
const webpack = require('webpack')
const {CleanWebpackPlugin} = require('clean-webpack-plugin')


let PAGE_SRC_PATH = path.resolve(__dirname, 'src/pages')
let BUILD_PATH = path.resolve(__dirname, 'build')
let DIST_PATH = path.resolve(__dirname, 'dist')
let files = glob.sync(PAGE_SRC_PATH + '/*/*.js')

let _bs = fs.stat(BUILD_PATH, function (err, s) {
    if (err != null) {
        fs.mkdir(BUILD_PATH, 0777, function () {
            console.log('mk build path success');
        });
    }
})
let _ds = fs.stat(DIST_PATH, function (err, s) {
    if (err != null) {
        fs.mkdir(DIST_PATH, 0777, function () {
            console.log('mk dist path success');
        });
    }
})

// console.log(files)
let filemap = {}
files.map(f => {
    let start = f.lastIndexOf('\/') + 1,
        end = f.lastIndexOf('.')
    let _name = f.substr(start, end - start)
    // console.log(start, end, _name);
    filemap[_name] = f;
})
console.log(filemap)
module.exports = {
    mode: 'development',
    entry: filemap,
    output: {
        path: BUILD_PATH,
        filename: '[name].js'
    },
    module: {
        rules: [{
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader'
                },
            },
            {
                test: /\.css$/,
                exclude: /node_modules/,
                use: {
                    loader: 'css-loader'
                }
            },
            {
                test: /\.(png|jpe?g|gif|svg)$/i,
                loader: 'file-loader',
                options: {
                    outputPath: 'images',
                }
            }
        ]
    },
    plugins: [
        new webpack.ProgressPlugin(),
        new HtmlwebpackPlugin({
            template: path.resolve(__dirname, 'public/index.html')
        }),
        new CleanWebpackPlugin({
            cleanOnceBeforeBuildPatterns: [path.resolve(__dirname, 'build')]
        })
    ],
    resolve: {
        extensions: ['.js', '.jsx'],
        modules: ['node_modules']
    }
}