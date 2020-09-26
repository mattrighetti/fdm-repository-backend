const express = require('express');
const router = express.Router()
const dataServices = require('../services/data.services');

module.exports = router;

// TODO changed endpoint name
router.get("/chart", function (req, res) {
    dataServices.getDataChart()
    .then(result => {
        res.json(result);
    });
});