-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE PROCEDURE insert_book(
    IN p_book_id UUID,
    IN p_book_data JSONB
)
LANGUAGE plpgsql
AS $$
DECLARE
    v_chapter_id UUID;
    v_book_data JSONB;
    v_chapter_data JSONB;
BEGIN
    FOR v_book_data IN SELECT * FROM jsonb_array_elements(p_book_data)
    LOOP
        INSERT INTO chapters (num_chapter, num_verses, book_id)
        VALUES (
            (v_book_data ->> 'numChapter')::INT,
            (v_book_data ->> 'totalVer')::INT,
            p_book_id
        ) RETURNING id INTO v_chapter_id;

        FOR v_chapter_data IN SELECT * FROM jsonb_array_elements((v_book_data ->> 'vers')::JSONB)
        LOOP
            INSERT INTO verses (num_verse, text, chapter_id)
            VALUES (
                (v_chapter_data ->> 'verNum')::INT,
                v_chapter_data ->> 'text',
                v_chapter_id
            );
        END LOOP;
    END LOOP;
EXCEPTION
    WHEN OTHERS THEN
        RAISE;
END;
$$;
-- +goose StatementEnd

-- +goose Down
DROP PROCEDURE insert_book;