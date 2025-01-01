cheers:
    @echo "Na zdraví! 🍺🍺🍺🍺🍺"

dev:
    @docker compose -f compose.yml up --build

run:
    @docker compose -f compose.yml up --build -d

stop:
    @docker compose -f compose.yml down

down:
    @docker compose -f compose.yml down --remove-orphans

rows:
    @sh -c 'find . -type d -name "src" | while IFS= read -r dir; do find "$dir" -type f -name "*.py" -exec cat {} +; done | wc -l'
