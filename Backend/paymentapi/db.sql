CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` varchar(255) DEFAULT NULL,
  `imageid` varchar(255) DEFAULT NULL,
  `paymentid` varchar(255) DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `created_on` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `payment_method_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB

insert into payment_methods(name, description) VALUES('stripe', 'Stripe Payment Gateway');
insert into orders (userid, imageid, paymentid, amount) VALUES('tim@gmail.com','982649','7646111188','50.00',1);
insert into orders (userid, imageid, paymentid, amount) VALUES('jane@gmail.com','672647','3546667788','22.50',1);
insert into orders (userid, imageid, paymentid, amount) VALUES('matt@gmail.com','672648','3546000088','50.00',1);
