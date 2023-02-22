module.exports = (mongoose: any) => {
  const schema = mongoose.Schema(
    {
      first_name: String,
      last_name: String,
      birth_date: Date,
      address: String,
    },
    {
      timestamps: true,
    }
  );

  return mongoose.model("users", schema);
};
