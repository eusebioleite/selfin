create table if not exists
    users (
        id integer primary key autoincrement,
        name text not null,
        login text not null,
        password text not null,
        image_url text not null,
        enabled integer not null default 0
    );

insert into
    users (name, login, password, image_url, enabled)
values
    ('admin', 'admin', 'admin', '', 1);

create table if not exists
    categories (
        id integer primary key autoincrement,
        description text not null
    );

insert into
    categories (description)
values
    ('Outros');

insert into
    categories (description)
values
    ('Receitas');

insert into
    categories (description)
values
    ('Tabaco');

insert into
    categories (description)
values
    ('Transporte');

insert into
    categories (description)
values
    ('Mercado');

insert into
    categories (description)
values
    ('Alimentação');

create table if not exists
    transactions (
        id integer primary key autoincrement,
        date text not null default (strftime('%d/%m/%Y', 'now')),
        amount int not null default 0,
        type text not null default 'Despesa',
        description text not null,
        category_id integer not null,
        user_id integer not null,
        foreign key (category_id) references categories (id),
        foreign key (user_id) references users (id)
    );

insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '01/06/2026',
        580,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '01/06/2026',
        14626,
        'Despesa',
        'Transferência enviada pelo Pix - NEOLINK - 10.749.772/0001-84 - COOP SICREDI DEXIS Agência: 718 Conta: 12664-5',
        1,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '01/06/2026',
        665,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '01/06/2026',
        598,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '02/06/2026',
        6280,
        'Despesa',
        'Transferência enviada pelo Pix - A. A. Q. Sardinha - 18.851.596/0001-07 - COOP SICREDI DEXIS Agência: 718 Conta: 24044-7',
        2,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '02/06/2026',
        720,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '02/06/2026',
        1647,
        'Despesa',
        'Transferência enviada pelo Pix - S O S LAR - 08.993.715/0001-02 - PAGSEGURO INTERNET IP S.A. (0290) Agência: 1 Conta: 53796210-2',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '02/06/2026',
        450,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '02/06/2026',
        450,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '03/06/2026',
        1647,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '04/06/2026',
        2457,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '05/06/2026',
        18990,
        'Despesa',
        'Transferência enviada pelo Pix - PAGSEGURO INTERNATIONAL - 06.375.668/0003-61 - PAGSEGURO INTERNET IP S.A. (0290) Agência: 1 Conta: 9671445-6',
        1,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '05/06/2026',
        590,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '05/06/2026',
        660,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '05/06/2026',
        2139,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '06/06/2026',
        2218,
        'Despesa',
        'Transferência enviada pelo Pix - S O S LAR - 08.993.715/0001-02 - PAGSEGURO INTERNET IP S.A. (0290) Agência: 1 Conta: 53796210-2',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '06/06/2026',
        670,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '07/06/2026',
        2127,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '07/06/2026',
        7400,
        'Despesa',
        'Transferência enviada pelo Pix - DAVID TENORIO ALECRIM MULTILIVROS - 57.263.619/0001-06 - MERCADO PAGO IP LTDA. (0323) Agência: 1 Conta: 6887225668-0',
        1,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '08/06/2026',
        800,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '08/06/2026',
        840,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '08/06/2026',
        7279,
        'Despesa',
        'Transferência enviada pelo Pix - A. A. Q. Sardinha - 18.851.596/0001-07 - COOP SICREDI DEXIS Agência: 718 Conta: 24044-7',
        2,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '08/06/2026',
        396,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '08/06/2026',
        1196,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '09/06/2026',
        2194,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '09/06/2026',
        300,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '09/06/2026',
        790,
        'Despesa',
        'Transferência enviada pelo Pix - MAYCON LUIS DOS SANTOS - •••.420.168-•• - BANCO INTER (0077) Agência: 1 Conta: 3668838-0',
        1,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '09/06/2026',
        840,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '09/06/2026',
        1010,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '10/06/2026',
        1196,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '10/06/2026',
        590,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '10/06/2026',
        660,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '10/06/2026',
        1047,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '11/06/2026',
        1794,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '11/06/2026',
        580,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '11/06/2026',
        920,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '11/06/2026',
        1196,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '12/06/2026',
        5000,
        'Receita',
        'Transferência Recebida - SILVANIA LEANDRO DA SILVA 29374156873 - 17.487.911/0001-98 - NU PAGAMENTOS - IP (0260) Agência: 1 Conta: 129289438-2',
        6,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '12/06/2026',
        1778,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '13/06/2026',
        1200,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '13/06/2026',
        1719,
        'Despesa',
        'Transferência enviada pelo Pix - S O S LAR - 08.993.715/0001-02 - PAGSEGURO INTERNET IP S.A. (0290) Agência: 1 Conta: 53796210-2',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '15/06/2026',
        39800,
        'Receita',
        'Transferência recebida pelo Pix - DOLPHIN TECNOLOGIA - 09.243.853/0001-29 - CORA SCFI (0403) Agência: 1 Conta: 3812647-2',
        6,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '15/06/2026',
        1578,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '15/06/2026',
        2120,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '15/06/2026',
        5080,
        'Despesa',
        'Compra no débito - FRANGO ASSADO PLANET',
        1,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        795,
        'Despesa',
        'Compra no débito - IFD*IFOOD CLUB',
        5,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        580,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        660,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        1678,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        790,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        4880,
        'Despesa',
        'Transferência enviada pelo Pix - A. A. Q. Sardinha - 18.851.596/0001-07 - COOP SICREDI DEXIS Agência: 718 Conta: 24044-7',
        2,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        920,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '16/06/2026',
        440,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '17/06/2026',
        580,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '17/06/2026',
        660,
        'Despesa',
        'Compra no débito via NuPay - 99',
        3,
        1
    );
insert into
    transactions (
        date,
        amount,
        type,
        description,
        category_id,
        user_id
    )
values
    (
        '17/06/2026',
        3018,
        'Despesa',
        'Compra no débito - SOSLar',
        4,
        1
    );