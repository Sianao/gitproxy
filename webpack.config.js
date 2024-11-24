const path = require('path');

module.exports = {
   
    entry: './src/pages/index.tsx', // 入口文件
    output: {
        path: path.resolve(__dirname, 'dist'), // 输出目录
        filename: 'bundle.js', // 输出文件名
    },
    module: {
        rules: [
            { test: /\.txt$/, use: 'raw-loader' },
            { test: /\.css$/, use: 'css-loader' },
            { test: /\.ts$/, use: 'ts-loader' },
            {
                test: /\.tsx?$/,
                use: {
                    loader: "babel-loader",
                    options: {
                      presets: [["@babel/preset-env", { targets: "defaults" }]],
                    },
          
                }
                
              },
        ]

    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
      },
    mode: 'development', // 根据需要设置模式
};

