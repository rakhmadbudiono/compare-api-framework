exports.up = function(knex) {
  return knex.schema.createTable("book", function(table) {
    table
      .increments("id_book")
      .primary()
      .notNullable()
      .unsigned();
    table.string("title").notNullable();
  });
};

exports.down = function(knex) {
  return knex.schema.dropTable("book");
};
