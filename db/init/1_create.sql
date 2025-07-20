\c app_db app

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  user_id            INTEGER      NOT NULL,
  user_name          VARCHAR(15)  NOT NULL,
  user_address       VARCHAR(100) NOT NULL,
  user_phone_number  CHAR(13),
  PRIMARY KEY (user_id)
);

INSERT INTO users VALUES
  (1, 'John Doe', '123 Main St, Springfield', '1234567890123'),
  (2, 'Jane Smith', '456 Elm St, Springfield', '9876543210987'),
  (3, 'Alice Johnson', '789 Oak St, Springfield', '1231231231234'),
  (4, 'Bob Brown', '321 Pine St, Springfield', '4564564564567');
