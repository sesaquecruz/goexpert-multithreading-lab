# Multithreading Lab - Go Expert

This project contains a simple Go program that demonstrates the use of multithreading through goroutines. The program concurrently calls two APIs, [ApiCep](https://apicep.com/api-de-consulta/) and [ViaCep](https://viacep.com.br/), with a provided cep code, and displays the response from the first API to return. If neither API returns a response within 1 second, the program will indicate a timeout was reached.

## Requirements

To use this program, you will need:

- Docker
- A stable internet connection

## Installation

1. Clone this repository:

```
git clone https://github.com/sesaquecruz/goexpert-multithreading-lab
```

2. Enter the project directory:

```
cd goexpert-multithreading-lab
```

3. Run the docker compose:

```
docker compose up --build
```

## Usage

1. In the project directory, enter the *getcep* container.

```
docker compose exec getcep sh
```

2. Inside the container, run the getcep:

```
./getcep <cep-number>
```

Replace `<cep-number>` with the desired cep code (ex: `./getcep 45520-000`).

## Troubleshooting

- If 1 second timeout is too short, consider changing it.

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) file for more information.
