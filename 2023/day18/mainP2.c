// https://github.com/LReg/clib
#include "../../clib/lib.c"
#define MAPWIDTH 300000

typedef enum Direction {
    UP,
    DOWN,
    LEFT,
    RIGHT
} Direction;
typedef enum LineType {
    HORIZONTAL,
    VERTICAL
} LineType;

typedef struct DigPart {
    size_t length;
    Direction dir;
    char color[6];
} DigPart;
typedef struct Point {
    int x, y;
} Point;

typedef struct Line {
    Point p1, p2;
    LineType type;
} Line;

DigPart * digParts[1000];
int digPartCount = 0;

Line * lines[1000];
int lineCount = 0;

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
        case '0':
            return RIGHT;
        case '1':
            return DOWN;
        case '2':
            return LEFT;
        case '3':
            return UP;
        default:
            assert(false);
    }
}

void printDigPart(DigPart * digPart) {
    printf("DigPart: %zu %d %s\n", digPart->length, digPart->dir, digPart->color);
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

/*
 Each hexadecimal code is six hexadecimal digits long.
 The first five hexadecimal digits encode the distance in meters as a five-digit hexadecimal number.
 The last hexadecimal digit encodes the direction to dig: 0 means R, 1 means D, 2 means L, and 3 means U.
 */
void fixInputs() {
    for (int i = 0; i < digPartCount; i++) {
        DigPart *digPart = digParts[i];
        digPart->dir = toDir(digPart->color[5]);
        char numberPart[5];
        strncpy(numberPart, digPart->color, 5);
        digPart->length = strtol(numberPart, NULL, 16);
        printDigPart(digPart);
    }
}

void generateLines() {
    Point pos = { 1000000, 1000000 };
    for (int i = 0; i < digPartCount; i++) {
        DigPart *digPart = digParts[i];
        Line * line = malloc(sizeof(Line));
        line->p1 = pos;
        switch (digPart->dir) {
            case UP:
                pos.y -= digPart->length;
                line->type = VERTICAL;
                break;
            case DOWN:
                pos.y += digPart->length;
                line->type = VERTICAL;
                break;
            case LEFT:
                pos.x -= digPart->length;
                line->type = HORIZONTAL;
                break;
            case RIGHT:
                pos.x += digPart->length;
                line->type = HORIZONTAL;
                break;
            default:
                assert(false);
        }
        line->p2 = pos;
        lines[lineCount++] = line;
    }
}

Line * isOnLine(Point p) {
    for (int i = 0; i < lineCount; i++) {
        Line * line = lines[i];
        if (line->type == HORIZONTAL) {
            if (p.y == line->p1.y && p.x >= line->p1.x && p.x <= line->p2.x)
                return lines[i];
        }
        else {
            if (p.x == line->p1.x && p.y >= line->p1.y && p.y <= line->p2.y)
                return lines[i];
        }
    }
    return NULL;
}

Point * unknownPoints[1000000];

size_t coutedCoveredTiles() {
    size_t count = 0;
    for (size_t y= 0; y < MAPWIDTH; y++) {
        bool inside = true;
        bool onHLine = false;
        for (size_t x= 0; x < MAPWIDTH; x++) {
            Point p = { x, y };
            Line * line = isOnLine(p);
            if (line == NULL && inside) {
                count++;
            }
            if (line != NULL) {
                if (line->type == HORIZONTAL) {
                    onHLine = true;
                    count++;
                }
                else {
                    inside = !inside;
                    count ++;
                }
            }
        }
    }
    return count;

}

int main(int argc, char *argv[]) {
    char *path = "day18/input.txt";
    if (1 || argc > 1 && strcmp(argv[1], "test") == 0)
        path = "day18/input.test.txt";
    FILE *input = assertOpenFile(path, "r");
    char lineBuffer[100];
    // clear map
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
    fixInputs();
    //printDigPartsToMap();
    printf("Covered Tiles: %zu\n", countCoveredTiles());
    //printMap();
}