module.exports = (app: any) => {
  const users = require("../controller/users.controller");
  const r = require("express").Router();

  r.post("/", users.create);

  app.use("/users", r);
};
