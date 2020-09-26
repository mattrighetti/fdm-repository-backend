const express = require('express');
const router = express.Router();
const missingModelsServices = require('../services/missing_models.services');

module.exports = router;

router.get("/", function (req, res) {
    missingModelsServices.getAllMissingModels()
    .then(result => {
        res.json(result);
    });
});

router.get("/count", function (req, res) {
    missingModelsServices.countMissingModels()
    .then(result => {
        res.json(result);
    });
});