# nature-remo-go

## Usage

```go
import natureremo "github.com/70-10/nature-remo-go"

func main() {
  c := natureremo.NewClient("<YOUR API TOKEN>")
  .
  .
  .
}
```

## CLI

### Install

```
go get -u github.com/70-10/nature-remo-go/cmd/nature-remo
```

## TODO

* [x] `GET /1/users/me`
* [ ] `POST /1/users/me`
* [x] `GET /1/devices`
* [ ] `POST /1/detectappliance`
* [x] `GET /1/appliances`
* [ ] `POST /1/appliances`
* [ ] `POST /1/appliance_orders`
* [ ] `POST /1/appliances/{appliance}/delete`
* [ ] `POST /1/appliances/{appliance}`
* [ ] `POST /1/appliances/{appliance}/aircon_settings`
* [x] `GET /1/appliances/{appliance}/signals`
* [ ] `POST /1/appliances/{appliance}/signals`
* [ ] `POST /1/appliances/{appliance}/signal_orders`
* [ ] `POST /1/signals/{signal}`
* [ ] `POST /1/signals/{signal}/delete`
* [x] `POST /1/signals/{signal}/send`
* [ ] `POST /1/devices/{device}`
* [ ] `POST /1/devices/{device}/delete`
* [ ] `POST /1/devices/{device}/temperature_offset`
* [ ] `POST /1/devices/{device}/humidity_offset`

## Documents

[API Document](http://swagger.nature.global)
