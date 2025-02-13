CREATE OR REPLACE FUNCTION generate_random_string(length INT) RETURNS TEXT AS $$
DECLARE
    chars TEXT := 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    result TEXT := '';
BEGIN
    LOOP
        result := '';
        FOR i IN 1..length LOOP
                result := result || substr(chars, floor(random() * length(chars))::int + 1, 1);
        END LOOP;
        IF NOT EXISTS (SELECT 1 FROM wallets WHERE address = result) THEN
            EXIT;
        END IF;
    END LOOP;

    RETURN result;
END;
$$ LANGUAGE plpgsql;

INSERT INTO wallets (address, balance)
SELECT generate_random_string(64), 100.0
FROM generate_series(1, 10)