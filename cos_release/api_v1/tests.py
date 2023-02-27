from django.contrib.auth.models import User, Permission
from django.test import TestCase

class MyTestCase(TestCase):
    def setUp(self):
        # Créez un utilisateur de test avec les permissions d'administrateur
        self.user = User.objects.create_user(
            username='admin',
            password='password',
            is_staff=True,
            is_superuser=True,
        )

        # Obtenez les permissions d'administrateur pour les tests
        self.ct_change_permission = Permission.objects.get(codename='change_mondelobjet')

        # Attribuez les permissions d'administrateur à l'utilisateur de test
        self.user.user_permissions.add(self.ct_change_permission)
        self.user.save()

    def test_my_function(self):
        # Utilisez l'utilisateur de test avec les permissions d'administrateur pour exécuter le test
        self.client.force_login(self.user)
        response = self.client.get('/mon_url/')
        self.assertEqual(response.status_code, 200)
