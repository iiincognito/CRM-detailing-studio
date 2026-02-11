сущность User
{
uuid
FIO
email
phone
birthday
}

сущность Service
{
id
name
price
}

сущность Admin
{
id
FIO
phone
}

сущность Order
{
id
user_id -- идентификатор пользователя
services_list -- набор услуг, как массив
created_date
additional_expenses
total price
created_by_admin_id
status -- ('new', 'in_progress', 'completed', 'cancelled', 'refunded')
}

Связующая таблица «многие-ко-многим» между заказами и услугами
{
order_id
service_id
price_at_order -- цена услуги на момент заказа (чтобы история не менялась при изменении цены услуги)
}