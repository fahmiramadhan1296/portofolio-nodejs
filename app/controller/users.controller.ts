const dbModel = require("../model");
const user = dbModel.users;

exports.create = (req: any, res: any) => {
  req.body.birth_date = new Date(req.body.birth_date);
  user
    .create(req.body)
    .then(() => res.send("Success"))
    .catch((err: any) => res.status(500).send({ message: err.message }));
};
