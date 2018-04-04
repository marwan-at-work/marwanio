const path = require('path');

module.exports = {
    entry: [
        './libs.js',
    ],
    output: {
        path: path.resolve(__dirname),
        filename: 'libs.inc.js',
    },
    mode: process.env.NODE_ENV || "development",
};