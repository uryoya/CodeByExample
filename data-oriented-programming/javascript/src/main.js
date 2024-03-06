const _ = require("lodash");

const watchmenBook = {
  isbn: "978-9863479116",
  title: "Watchmen",
  publicationYear: 1987,
  authorIds: ["alan-moore", "dave-gibbons"],
  bookItems: [
    {
      id: "book-item-1",
      libId: "nyc-central-lib",
      isLent: true,
    },
    {
      id: "book-item-2",
      libId: "nyc-central-lib",
      isLent: false,
    },
  ],
};

const catalogData = {
  booksByIsbn: {
    "978-9863479116": watchmenBook,
  },
  authorsById: {
    "alan-moore": {
      name: "Alan Moore",
      booksIsbns: ["978-9863479116"],
    },
    "dave-gibbons": {
      name: "Dave Gibbons",
      booksIsbns: ["978-9863479116"],
    },
  },
};

const libraryData = {
  catalog: catalogData,
};

class Catalog {
  static authorNames(catalogData, book) {
    const authorIds = _.get(book, "authorIds");
    const names = _.map(authorIds, (authorId) =>
      _.get(catalogData, ["authorsById", authorId, "name"])
    );
    return names;
  }

  static bookInfo(catalogData, book) {
    const bookInfo = {
      title: _.get(book, "title"),
      isbn: _.get(book, "isbn"),
      authorNames: Catalog.authorNames(catalogData, book),
    };
    return bookInfo;
  }

  static searchBooksByTitle(catalogData, query) {
    const allBooks = _.get(catalogData, "booksByIsbn");
    const matchingBooks = _.filter(allBooks, (book) =>
      _.get(book, "title").includes(query)
    );
    const bookInfos = _.map(matchingBooks, (book) =>
      Catalog.bookInfo(catalogData, book)
    );
    return bookInfos;
  }
}

class Library {
  static searchBooksByTitleJSON(libraryData, query) {
    const catalogData = _.get(libraryData, "catalog");
    const results = Catalog.searchBooksByTitle(catalogData, query);
    return JSON.stringify(results);
  }
}

var result = Library.searchBooksByTitleJSON(libraryData, "Wat");

console.log(result);
