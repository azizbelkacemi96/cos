import psycopg2

try:
    connection = psycopg2.connect(user="db_user",
                                  password="db_password",
                                  host="db_host",
                                  port="db_port",
                                  database="tower_database")

    cursor = connection.cursor()

    # Compte le nombre d'hôtes dans la table main_host
    cursor.execute("SELECT COUNT(*) FROM main_host;")

    host_count = cursor.fetchone()[0]

    print(f"Total hosts: {host_count}")

except (Exception, psycopg2.Error) as error:
    print("Error while connecting to PostgreSQL", error)

finally:
    if connection:
        cursor.close()
        connection.close()
