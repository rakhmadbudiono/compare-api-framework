const app = require("express")();
const bodyParser = require("body-parser");

const book = require("./book");

const PORT = process.env.PORT || 3000;

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get("/", async (req, res) => {
  const result = await book.getBooks();

  res.json(result);
});

app.post("/", async (req, res) => {
  const result = await book.createBook(req.body);

  res.end();
});

app.get("/:id", async (req, res) => {
  const id = parseInt(req.params.id);
  const result = await book.getBookById(id);

  res.json(result);
});

app.put("/:id", async (req, res) => {
  const id = parseInt(req.params.id);
  await book.updateBook(id, req.body);

  res.status(204).end();
});

app.delete("/:id", async (req, res) => {
  const id = parseInt(req.params.id);
  await book.deleteBook(id);

  res.status(204).end();
});

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}.`);
  console.log(`Started at ${new Date().toUTCString()}`);
});
