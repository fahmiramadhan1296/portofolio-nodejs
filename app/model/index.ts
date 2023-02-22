const dbConfig = require("../config/database");
const mongoose = require("mongoose");

module.exports = {
  mongoose,
  url: dbConfig.urlMongoProd,
  users: require("./users.model.ts")(mongoose),
};
