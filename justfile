cheers:
    @echo "Na zdraví! 🍺🍺🍺🍺🍺"

dev:
    @docker compose -f docker/compose.yml up --build

run:
    @docker compose -f docker/compose.yml up --build -d

stop:
    @docker compose -f docker/compose.yml down

down:
    @docker compose -f docker/compose.yml down --remove-orphans

tests:
    @echo "IMPLEMENT IT"

rows:
    @sh -c 'find . -type d -name "src" | while IFS= read -r dir; do find "$dir" -type f -name "*.py" -exec cat {} +; done | wc -l'
