FROM cassandra:4

RUN sed -i -r 's/authenticator: AllowAllAuthenticator/authenticator: PasswordAuthenticator/' /etc/cassandra/cassandra.yaml
# RUN sed -i -r 's/user_defined_functions_enabled: false/user_defined_functions_enabled: true/' /etc/cassandra/cassandra.yaml

CMD ["cassandra", "-f"]
