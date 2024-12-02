// https://github.com/LReg/clib
#include "../../clib/lib.c"
#define MAPWIDTH 1000

typedef enum Direction {
    UP,
    DOWN,
    LEFT,
    RIGHT
} Direction;
typedef struct DigPart {
    int length;
    Direction dir;
    char color[6];
} DigPart;
typedef struct Point {
    int x, y;
} Point;

DigPart * digParts[1000];
char map[MAPWIDTH][MAPWIDTH];
int digPartCount = 0;

Direction toDir(char c) {
    switch (c) {
        case 'U':
            return UP;
        case 'D':
            return DOWN;
        case 'L':
            return LEFT;
        case 'R':
            return RIGHT;
        default:
            assert(false);
    }
}

void printDigPart(DigPart * digPart) {
    printf("DigPart: %d %d %s\n", digPart->length, digPart->dir, digPart->color);
}

void printDigPartsToMap() {
    int x = MAPWIDTH / 2, y = MAPWIDTH / 2;
    Direction dir;
    for (int i = 0; i < digPartCount; i++) {
        DigPart *digPart = digParts[i];
        dir = digPart->dir;
        for (int j = 0; j < digPart->length; j++) {
            switch (dir) {
                case UP:
                    y--;
                    break;
                case DOWN:
                    y++;
                    break;
                case LEFT:
                    x--;
                    break;
                case RIGHT:
                    x++;
                    break;
                default:
                    assert(false);
            }
            // check out of map
            if (x < 0 || x >= MAPWIDTH || y < 0 || y >= MAPWIDTH)
                assert(false);
            //printf("x: %d, y: %d\n", x, y);
            map[y][x] = '#';
        }
    }
}

void printMap() {
    for (int y = 0; y < MAPWIDTH; y++) {
        for (int x = 0; x < MAPWIDTH; x++) {
            putchar(map[y][x]);
        }
        putchar('\n');
    }
}

Point findFloodfillPoint() {
    for (int y = 0; y < MAPWIDTH; y++) {
        bool onEdge = false;
        bool inside = false;
        for (int x = 0; x < MAPWIDTH; x++) {
            if (!onEdge && map[y][x] == '#') {
                onEdge = true;
            }
            else if (onEdge && map[y][x] == '.') {
                inside = true;
            }
            else if (onEdge && map[y][x] == '#') {
                inside = false;
                onEdge = false;
                break;
            }
            if (inside) {
                return (Point) { x, y };
            }
        }
    }
}

size_t countCoveredTiles() {
    size_t count = 0;
    Point floodfillPoint = findFloodfillPoint();
    map[floodfillPoint.y][floodfillPoint.x] = 'O';
    count++;
    bool done = false;
    while (!done) {
        done = true;
        for (int y = 0; y < MAPWIDTH; y++) {
            for (int x = 0; x < MAPWIDTH; x++) {
                if (map[y][x] == 'O') {
                    if (map[y - 1][x] == '.') {
                        map[y - 1][x] = 'O';
                        count++;
                        done = false;
                    }
                    if (map[y + 1][x] == '.') {
                        map[y + 1][x] = 'O';
                        count++;
                        done = false;
                    }
                    if (map[y][x - 1] == '.') {
                        map[y][x - 1] = 'O';
                        done = false;
                        count++;
                    }
                    if (map[y][x + 1] == '.') {
                        map[y][x + 1] = 'O';
                        done = false;
                        count++;
                    }
                }
            }
        }
    }
    for (int y = 0; y < MAPWIDTH; y++)
        for (int x = 0; x < MAPWIDTH; x++)
            if (map[y][x] == '#')
                count++;
    return count;
}

int main(int argc, char *argv[]) {
    char *path = "day18/input.txt";
    if (argc > 1 && strcmp(argv[1], "test") == 0)
        path = "day18/input.test.txt";
    FILE *input = assertOpenFile(path, "r");
    char lineBuffer[100];
    // clear map
    memset(map, '.', sizeof(map));
    while(fgets(lineBuffer, 100, input) != NULL) {
        DigPart *digPart = malloc(sizeof(DigPart));
        digPart->dir = toDir(lineBuffer[0]);
        char * len = skipNonDigits(lineBuffer);
        digPart->length = atoi(len);
        char * color = strchr(skipDecimals(len), '#') + 1;
        if (color != NULL) {
            strncpy(digPart->color, color, 6);
        }
        digParts[digPartCount++] = digPart;
    }
    printDigPartsToMap();
    printf("Covered Tiles: %zu\n", countCoveredTiles());
    //printMap();
}