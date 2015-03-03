CREATE VIEW
    user_fts_view AS
    (
        SELECT
            id AS id, to_tsvector(user_name)||to_tsvector(first_name)|| to_tsvector(last_name)||
            to_tsvector(email)|| to_tsvector(cell_phone)|| to_tsvector(office_phone)|| to_tsvector(fax)||
            to_tsvector(country)|| to_tsvector(city)|| to_tsvector(post_code)|| to_tsvector(company_name) as details
        FROM
            user_info
    );