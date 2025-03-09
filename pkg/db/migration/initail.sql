create table roles
(
    id          uuid default gen_random_uuid() not null
        primary key,
    name        varchar(50)                    not null
        unique,
    description text,
    code        varchar(50)                    not null
        unique
);

alter table roles
    owner to hotel_reservation;

create table permissions
(
    id          uuid default gen_random_uuid() not null
        primary key,
    name        varchar(100)                   not null
        unique,
    description text
);

alter table permissions
    owner to hotel_reservation;

create table role_permissions
(
    role_id       uuid not null
        references roles
            on delete cascade,
    permission_id uuid not null
        references permissions
            on delete cascade,
    primary key (role_id, permission_id)
);

alter table role_permissions
    owner to hotel_reservation;

create table users
(
    id            uuid      default gen_random_uuid() not null
        primary key,
    username      varchar(50)                         not null
        unique,
    email         varchar(100)                        not null
        unique,
    password_hash text                                not null,
    role_id       uuid
                                                      references roles
                                                          on delete set null,
    created_at    timestamp default now()
);

alter table users
    owner to hotel_reservation;

create table guests
(
    id         uuid      default gen_random_uuid() not null
        primary key,
    first_name varchar(50)                         not null,
    last_name  varchar(50)                         not null,
    email      varchar(100)                        not null
        unique,
    phone      varchar(20)                         not null
        unique,
    created_at timestamp default now()
);

alter table guests
    owner to hotel_reservation;

create table rooms
(
    id         uuid      default gen_random_uuid() not null
        primary key,
    number     varchar(10)                         not null
        unique,
    type       varchar(50)                         not null,
    price      numeric(10, 2)                      not null,
    status     varchar(20)                         not null
        constraint rooms_status_check
            check ((status)::text = ANY
        ((ARRAY ['Available'::character varying, 'Booked'::character varying, 'Under Maintenance'::character varying])::text[])),
    created_at timestamp default now()
);

alter table rooms
    owner to hotel_reservation;

create table reservations
(
    id         uuid      default gen_random_uuid() not null
        primary key,
    guest_id   uuid
        references guests
            on delete cascade,
    check_in   date                                not null,
    check_out  date                                not null,
    status     varchar(20)                         not null
        constraint reservations_status_check
            check ((status)::text = ANY
        ((ARRAY ['Confirmed'::character varying, 'Canceled'::character varying, 'Completed'::character varying])::text[])),
    created_at timestamp default now()
);

alter table reservations
    owner to hotel_reservation;

create table reservation_details
(
    id             uuid default gen_random_uuid() not null
        primary key,
    reservation_id uuid
        references reservations
            on delete cascade,
    room_id        uuid
        references rooms
            on delete cascade
);

alter table reservation_details
    owner to hotel_reservation;

create table payments
(
    id             uuid      default gen_random_uuid() not null
        primary key,
    reservation_id uuid
        references reservations
            on delete cascade,
    amount         numeric(10, 2)                      not null,
    payment_method varchar(50)                         not null
        constraint payments_payment_method_check
            check ((payment_method)::text = ANY
        ((ARRAY ['Credit Card'::character varying, 'Cash'::character varying, 'Bank Transfer'::character varying])::text[])),
    status         varchar(20)                         not null
        constraint payments_status_check
            check ((status)::text = ANY
                   ((ARRAY ['Paid'::character varying, 'Pending'::character varying, 'Failed'::character varying])::text[])),
    created_at     timestamp default now()
);

alter table payments
    owner to hotel_reservation;

create table invoices
(
    id           uuid default gen_random_uuid() not null
        primary key,
    payment_id   uuid
        references payments
            on delete cascade,
    total_amount numeric(10, 2)                 not null,
    issued_date  date                           not null
);

alter table invoices
    owner to hotel_reservation;

create table room_services
(
    id             uuid      default gen_random_uuid() not null
        primary key,
    reservation_id uuid
        references reservations
            on delete cascade,
    service_name   varchar(100)                        not null,
    cost           numeric(10, 2)                      not null,
    created_at     timestamp default now()
);

alter table room_services
    owner to hotel_reservation;

create table housekeeping
(
    id          uuid default gen_random_uuid() not null
        primary key,
    room_id     uuid
        references rooms
            on delete cascade,
    assigned_to uuid
                                               references users
                                                   on delete set null,
    cleaned_at  timestamp,
    status      varchar(20)                    not null
        constraint housekeeping_status_check
            check ((status)::text = ANY ((ARRAY ['Pending'::character varying, 'Completed'::character varying])::text[]))
    );

alter table housekeeping
    owner to hotel_reservation;

create table audit_logs
(
    id         uuid      default gen_random_uuid() not null
        primary key,
    user_id    uuid
                                                   references users
                                                       on delete set null,
    action     text                                not null,
    table_name varchar(50)                         not null,
    record_id  uuid                                not null,
    timestamp  timestamp default now()
);

alter table audit_logs
    owner to hotel_reservation;

CREATE TABLE permissions (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             name VARCHAR(100) UNIQUE NOT NULL,
                             description TEXT
);

alter table permissions
    owner to hotel_reservation;

CREATE TABLE role_permissions (
                                  role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
                                  permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
                                  PRIMARY KEY (role_id, permission_id)
);

alter table role_permissions
    owner to hotel_reservation;

ALTER TABLE users ADD COLUMN updated_at TIMESTAMP DEFAULT now();
ALTER TABLE rooms ADD COLUMN updated_at TIMESTAMP DEFAULT now();
ALTER TABLE reservations ADD COLUMN updated_at TIMESTAMP DEFAULT now();
ALTER TABLE payments ADD COLUMN updated_at TIMESTAMP DEFAULT now();
ALTER TABLE users ADD COLUMN active BOOLEAN DEFAULT TRUE;
ALTER TABLE payments ALTER COLUMN status TYPE TEXT USING status::TEXT;
ALTER TABLE payments ADD CONSTRAINT payments_status_enum CHECK (
    status IN ('Paid', 'Pending', 'Failed', 'Refunded')
    );

CREATE TABLE room_types (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name VARCHAR(50) UNIQUE NOT NULL,
                            description TEXT
);
ALTER TABLE rooms ADD COLUMN type_id UUID REFERENCES room_types(id);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_reservations_guest ON reservations(guest_id);
CREATE INDEX idx_payments_reservation ON payments(reservation_id);
CREATE INDEX idx_housekeeping_room ON housekeeping(room_id);


INSERT INTO roles (name, description, code)
VALUES
    ( 'Admin', 'Administrator with full access', 'ADMIN'),
    ( 'Receptionist', 'Handles guest reservations', 'RECEPTIONIST'),
    ( 'Housekeeping', 'Manages room cleaning status', 'HOUSEKEEPING');

INSERT INTO permissions ( name, description)
VALUES
    ( 'manage_users', 'Create, update, and delete users'),
    ( 'manage_reservations', 'Create, update, and delete reservations'),
    ( 'manage_payments', 'Handle all payment operations'),
    ( 'view_reports', 'Access financial and audit reports'),
    ( 'manage_housekeeping', 'Assign and update housekeeping tasks');