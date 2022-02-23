-- https://www.timescale.com/blog/select-the-most-recent-record-of-many-items-with-postgresql/

CREATE TABLE IF NOT EXISTS trucks
(
    truck_id      int PRIMARY KEY,
    make          varchar(255),
    model         varchar(255),
    weight_class  varchar(255),
    date_acquired bigint NOT NULL,
    active_status bool
);

CREATE TABLE IF NOT EXISTS truck_reading
(
    ts        bigint NOT NULL,
    truck_id  int,
    milage    int,
    fuel      int,
    latitude  float8,
    longitude float8
);
CREATE INDEX ix_ts ON truck_reading (ts DESC);
CREATE INDEX ix_truck_id_ts ON truck_reading (truck_id, ts DESC);

/*
 * The logging table alternative. The PRIMARY KEY will create an
*  index on the truck_id column to make querying for specific trucks more efficient
 */
CREATE TABLE IF NOT EXISTS truck_log
(
    truck_id  int PRIMARY KEY REFERENCES trucks (truck_id),
    last_time timestamp,
    milage    int,
    fuel      int,
    latitude  float8,
    longitude float8
);

/*
* Because the table will mostly be UPDATE heavy, a slightly reduced
* FILLFACTOR can alleviate maintenance contention and reduce
* page bloat on the table.
*/
ALTER TABLE truck_log
    SET (fillfactor =90);

/*
 * This is the trigger function which will be executed for each row
*  of an INSERT or UPDATE. Again, YMMV, so test and adjust appropriately
 */
CREATE OR REPLACE FUNCTION create_truck_trigger_fn()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL AS
$BODY$
BEGIN
    INSERT INTO truck_log
    VALUES (NEW.truck_id, NEW.time, NEW.milage, NEW.fuel, NEW.latitude, NEW.longitude)
    ON CONFLICT (truck_id) DO UPDATE SET last_time=NEW.time,
                                         milage=NEW.milage,
                                         fuel=NEW.fuel,
                                         latitude=NEW.latitude,
                                         longitude=NEW.longitude;
    RETURN NEW;
END
$BODY$;

/*
*  With the trigger function created, actually assign it to the truck_reading
*  table so that it will execute for each row
*/
CREATE TRIGGER create_truck_trigger
    BEFORE INSERT OR UPDATE
    ON truck_reading
    FOR EACH ROW
EXECUTE PROCEDURE create_truck_trigger_fn();
