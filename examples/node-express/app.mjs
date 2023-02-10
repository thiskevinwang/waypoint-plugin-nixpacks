import express from "express";
import logger from "morgan";

import indexRouter from "./routes/index.mjs";

const app = express();

app.use(logger("dev"));
app.use("/", indexRouter);

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  res.status(404).json({ message: "not found" });
});

// error handler
app.use(function (err, req, res, next) {
  res.status(500).json({ message: "internal server error" });
});

const PORT = 3000;

app.listen(PORT, () => {
  console.log("listening on port %d", PORT);
});
