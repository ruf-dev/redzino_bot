-- +goose Up
-- +goose StatementBegin
CREATE TABLE settings
(
    roll_cost          INT NOT NULL DEFAULT -5,
    roll_fruit_prize   INT NOT NULL DEFAULT 50,
    roll_jackpot_prize INT NOT NULL DEFAULT 150,
    dice_cost          INT NOT NULL DEFAULT -2,
    dice_win           INT NOT NULL DEFAULT 12
);

INSERT INTO settings
VALUES (-5, 50, 150, -2, 12);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
