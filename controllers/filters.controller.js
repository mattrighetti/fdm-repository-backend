const express = require('express');
const router = express.Router();

const filtersServices = require('../services/filters.services');

module.exports = router;

router.get("/eis", function (req, res) {
    filtersServices.getFiltersWithoutSlash()
    .then(result => {
        res.json(result);
    });
});

router.get("/country", function (req, res) {
    filtersServices.getCountriesOfDevelopment()
    .then(result => {
        res.json(result);
    });
});