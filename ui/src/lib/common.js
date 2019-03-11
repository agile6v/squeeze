
function getAppRoot() {
    if (process.env.NODE_ENV !== 'production') {
        return 'http://localhost:3000'
    }
    return ''
}

module.exports = {
    getAppRoot,
}