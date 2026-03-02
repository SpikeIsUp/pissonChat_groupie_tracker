const express = require('express');
const router = express.Router();
const sneakerController = require('../controller/sneaker');

router.get('/sneakers', sneakerController.getSneakers);
router.get('/sneaker/:id', sneakerController.getSneakerById);

module.exports = router;