const data = require("../data.json");

exports.getSneakers = (req, res) => {
  const sneakers = data.sneakers;
  res.status(200).json({
    code: 200,
    message: "Sneakers retrieved successfully",
    sneakers: sneakers,
  });
};

exports.getSneakerById = async (req, res) => {
  const id = req.params.id;
  const sneakers = data.sneakers;

  const sneaker = sneakers.find((sneaker) => sneaker.id === parseInt(id));
  if (!sneaker) {
    res.status(404).json({
      code: 404,
      message: "Sneaker not found",
    });
  } else {
    res.status(200).json({
      code: 200,
      message: "Sneaker retrieved successfully",
      sneaker: sneaker,
    });
  }
};
