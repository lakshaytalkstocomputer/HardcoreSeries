module.exports = {
    mode: 'development',
    entry: './app/app.ts',
    module:{
        rules:[
            {
                test:/\.tsx?$/,
                use:"ts-loader",
                exclude: /node_modules/,
            },
        ],
    },
    devtool: 'inline-source-map',
    resolve: {
        extensions : ['.tsx', ".ts", ".js"]
    },
    output : {
        filename : 'bundle.js',
    },
    devServer: {
        inline: true
    }
};