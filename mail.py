import smtplib
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from email.mime.base import MIMEBase
from email import encoders

def send_email_with_attachment(sender_email, receiver_email, subject, body, attached_file):
    # Create an MIMEMultipart object
    message = MIMEMultipart()
    message['From'] = sender_email
    message['To'] = receiver_email
    message['Subject'] = subject

    # Message body
    message.attach(MIMEText(body, 'plain'))

    # Open the file in binary mode
    with open(attached_file, 'rb') as attachment:
        part = MIMEBase('application', 'octet-stream')
        part.set_payload(attachment.read())

    # Encode the file in base64
    encoders.encode_base64(part)

    # Add file attachment headers
    part.add_header('Content-Disposition', f"attachment; filename= {attached_file}")

    # Attach the file to the message
    message.attach(part)

    # Establish an SMTP connection with the server
    smtp_server = smtplib.SMTP('smtp.example.com', 587)  # Replace with the appropriate SMTP server details
    smtp_server.starttls()

    # Log in to the sender's account
    sender_password = "your_password"  # Replace with the sender's password
    smtp_server.login(sender_email, sender_password)

    # Send the email
    full_message = message.as_string()
    smtp_server.sendmail(sender_email, receiver_email, full_message)

    # Quit the SMTP session
    smtp_server.quit()

    print("The email has been sent successfully.")

# Example usage of the function
sender_email = "your_email@gmail.com"
receiver_email = "recipient@example.com"
subject = "Subject of the email"
body = "This is the body of the email."
attached_file = "path/to/your_file.pdf"

send_email_with_attachment(sender_email, receiver_email, subject, body, attached_file)
