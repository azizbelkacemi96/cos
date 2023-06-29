import psycopg2

# Remplacez ces variables par vos informations de connexion
DB_NAME = "your_db_name"
DB_USER = "your_db_user"
DB_PASSWORD = "your_db_password"
DB_HOST = "your_db_host"
DB_PORT = "your_db_port"

try:
    # Établir la connexion à la base de données
    connection = psycopg2.connect(
        dbname=DB_NAME,
        user=DB_USER,
        password=DB_PASSWORD,
        host=DB_HOST,
        port=DB_PORT
    )
    
    # Créer un curseur
    cursor = connection.cursor()

    # Exécuter la requête SQL pour récupérer les hôtes
    cursor.execute("SELECT name FROM main_host;")
    hosts = cursor.fetchall()

    # Mettre les noms des hôtes en majuscules et supprimer les doublons
    hosts = list(set([host[0].upper() for host in hosts]))

    # Afficher les hôtes
    for host in hosts:
        print(host)

    # Fermer la connexion
    cursor.close()
    connection.close()

except Exception as error:
    print("Erreur lors de la connexion à la base de données:", error)
