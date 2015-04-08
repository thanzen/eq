exports.calculatePage = function (rowsPerPage, total) {
    if (total <= 0)
        return 1;
    if (rowsPerPage <= 0)
        throw "Nubmer must be great than 0!";
    return Math.ceil(total / rowsPerPage);
};
