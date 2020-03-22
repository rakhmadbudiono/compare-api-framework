const config = require("./knexfile");
const knex = require("knex")(config);

async function createBook(data) {
  result = await knex("book").insert({
    title: data.title
  });

  return result;
}

async function getBooks() {
  const result = await knex("book").select("*");

  return result;
}

async function getBookById(id) {
  const result = await knex("book")
    .select("*")
    .where("id_book", id);

  return result;
}

async function updateBook(id, data) {
  const result = await knex("book")
    .where("id_book", id)
    .update("title", data.title);

  return result;
}

async function deleteBook(id) {
  const result = await knex("book")
    .where("id_book", id)
    .del();

  return result;
}

module.exports = {
  createBook,
  getBooks,
  getBookById,
  updateBook,
  deleteBook
};
