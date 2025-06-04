from datetime import datetime
import argparse
import psycopg2

# Database connection configuration
DB_CONFIG = {
    'dbname': 'indicator_test',
    'user': 'pchen',
    'password': 'fZwHFgf*-T6Y',
    'host': '127.0.0.1',
    'port': '5432'
}

# SQL for creating the table (if it doesn't exist)
CREATE_TABLE_SQL = """
CREATE TABLE IF NOT EXISTS table_test (
    id SERIAL PRIMARY KEY,
    action VARCHAR(50) NOT NULL,
    step_name VARCHAR(50) NOT NULL,
    order_id BIGINT NOT NULL,
    data_time DOUBLE PRECISION NOT NULL,
    log_update_time TIMESTAMP NOT NULL,
    CONSTRAINT unique_log_entry UNIQUE (action, step_name, order_id, data_time)
);
"""

# SQL for inserting data
INSERT_SQL = """
INSERT INTO table_test3 (action, step_name, order_id, data_time, log_update_time)
VALUES (%s, %s, %s, %s, %s)
ON CONFLICT ON CONSTRAINT unique_log_entry3 DO NOTHING;
"""


def create_table(conn):
    """Create the table if it doesn't exist"""
    with conn.cursor() as cursor:
        cursor.execute(CREATE_TABLE_SQL)
        conn.commit()
        print("Table created or already exists")


def parse_log_line(line):
    """Parse a single line from the log file"""
    try:
        # Split by comma
        parts = line.strip().split(',', 4)

        if len(parts) != 5:
            print(f"Warning: Invalid line format: {line}")
            return None

        action = parts[0].strip()
        step_name = parts[1].strip()
        order_id = parts[2].strip()
        data_time = parts[3].strip()
        log_update_time = parts[4].strip()

        # Convert to appropriate data types
        order_id = int(order_id)
        data_time = float(data_time)
        log_update_time = datetime.strptime(log_update_time, '%Y-%m-%d %H:%M:%S.%f')

        return (action, step_name, order_id, data_time, log_update_time)

    except Exception as e:
        print(f"Error parsing line: {line}")
        print(f"Exception: {e}")
        return None


def import_log_data(conn, log_file_path, batch_size=100):
    """Import data from log file into database using batch processing"""
    inserted_count = 0
    error_count = 0
    batch_count = 0
    batch = []

    with conn.cursor() as cursor:
        with open(log_file_path, 'r') as file:
            for line in file:
                if not line.strip():
                    continue

                data = parse_log_line(line)
                if data:
                    batch.append(data)

                    # When batch is full, execute and commit
                    if len(batch) >= batch_size:
                        batch_count += 1
                        print(f"Processing batch #{batch_count} ({len(batch)} records)...")

                        for record in batch:
                            try:
                                cursor.execute(INSERT_SQL, record)
                                inserted_count += 1
                            except Exception as e:
                                print(f"Error inserting data: {record}")
                                print(f"Exception: {e}")
                                error_count += 1

                        conn.commit()
                        batch = []  # Clear the batch

            # Process any remaining records
            if batch:
                batch_count += 1
                print(f"Processing final batch #{batch_count} ({len(batch)} records)...")

                for record in batch:
                    try:
                        cursor.execute(INSERT_SQL, record)
                        inserted_count += 1
                    except Exception as e:
                        print(f"Error inserting data: {record}")
                        print(f"Exception: {e}")
                        error_count += 1

                conn.commit()

    print(f"Import completed. Batches: {batch_count}, Inserted: {inserted_count}, Errors: {error_count}")


def main(log_file_path: str):
    print(f"Using log file: {log_file_path}")
    # Connect to the database
    conn = None
    try:
        conn = psycopg2.connect(**DB_CONFIG)
        print("Connected to PostgreSQL database")
        conn.autocommit = True
        # Create table if it doesn't exist
        # create_table(conn)

        # Import log data
        # import_log_data(conn, log_file_path)
        # Import log data with batch processing
        batch_size = 100
        import_log_data(conn, log_file_path, batch_size)

    except Exception as e:
        print(f"Database error: {e}")
    finally:
        if conn:
            conn.close()
            print("Database connection closed")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Process log file path")
    parser.add_argument("log_file_path", type=str, help="Path to the log file")
    args = parser.parse_args()
    main(args.log_file_path)
