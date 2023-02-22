const express = require("express");
const cors = require("cors");
const db = require("./app/model");

const app = express();
const corsOptions = {
  origin: "*",
};

// register cors middleware
app.use(cors(corsOptions));
app.use(express.json());

// Connect to Database

const mongooseConfig = {
  useNewUrlParser: true,
  useUnifiedTopology: true,
};

db.mongoose
  .connect(db.url, mongooseConfig)
  .then((res: any) => console.log(res))
  .catch((err: any) => console.log(`Error in DB connection ${err}`));

// routes
require("./app/routes/users.routes")(app);

const PORT = process.env.PORT || 8000;
app.listen(PORT, () => console.log("Running"));
