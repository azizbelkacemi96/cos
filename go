---
- hosts: votre_hote
  gather_facts: no
  become: yes
  vars:
    repository_url: 'url_de_votre_repository'
    dest_directory: '/applis/26482-tdgg'
    owner: 'xlrel'
    group: 'xlrel'
    python_version: '3.8'
    service_name_nginx: 'nginx'
    service_name_gunicorn: 'gunicorn'

  tasks:
    - name: Cloner le git repository
      git:
        repo: "{{ repository_url }}"
        dest: "{{ dest_directory }}"
        clone: yes
        update: yes

    - name: Changer le propriétaire du répertoire
      file:
        path: "{{ dest_directory }}"
        owner: "{{ owner }}"
        group: "{{ group }}"
        state: directory
        recurse: yes

    - name: Créer un environnement virtuel Python
      command: python{{ python_version }} -m venv venv
      args:
        chdir: "{{ dest_directory }}"
        creates: "{{ dest_directory }}/venv"

    - name: Installer les dépendances à partir de requirements.txt
      pip:
        requirements: "{{ dest_directory }}/requirements.txt"
        virtualenv: "{{ dest_directory }}/venv"

    - name: Démarrer le service Nginx
      systemd:
        name: "{{ service_name_nginx }}"
        state: started
        enabled: yes

    - name: Démarrer le service Gunicorn
      systemd:
        name: "{{ service_name_gunicorn }}"
        state: started
        enabled: yes
...
