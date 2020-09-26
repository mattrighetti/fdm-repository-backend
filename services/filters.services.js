const db = require('./database.service');

module.exports = { 
    getFiltersWithoutSlash,
    getCountriesOfDevelopment
}

async function getFiltersWithoutSlash() {
    return db("models")
    .distinct('eis')
    .map(obj => obj.eis)
    .filter(obj => { return !obj.includes('/') })
}

async function getCountriesOfDevelopment() {
    return db("models")
    .distinct('cod')
    .map(obj => obj.cod)
}