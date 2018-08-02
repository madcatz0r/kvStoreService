# kvStoreService

Заготовка gRPC сервиса взаимодействия с key-value хранилищем

### Хранилище
На данный момент in-memory.
map[string]interface{} + RWMutex при количестве доступных потоков <=4,
sync.Map - >4

###TODO:
* Опробовать кастомный кодек
* Эксперимент с google/protobuf/any.proto
* Перевести на map[interface{}]interface{}
* Покрыть юнит-тестами + бенчмарки + комментарии
* Перевести этот поток сознания на международный