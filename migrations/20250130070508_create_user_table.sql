-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" 
(
    "id" varchar(63) PRIMARY KEY,
    "email" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "username" varchar(63) NOT NULL,
    "fullname" varchar(255),
    "phone_number" varchar(255),
    "modified_by" varchar(255) NOT NULL,
    "created_date" date NOT NULL,
    "modified_date" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";
-- +goose StatementEnd
