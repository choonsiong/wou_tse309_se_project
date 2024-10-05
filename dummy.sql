INSERT INTO "public"."publishers"("publisher_name")
VALUES ('No Starch Press'),
       ('SAP Press'),
       ('旗標科技'),
       ('Gumroad'),
       ('Pragmatic Programmer');

INSERT INTO publishers (publisher_name)
VALUES ('Unknown');


INSERT INTO mybooks (title, publisher_id, publication_year, slug, description)
VALUES ('The Shining', 8, 1977, 'the-shining',
        E'Jack Torrance, his wife Wendy, and their young son Danny move into the Overlook Hotel, where Jack has been hired as the winter caretaker. Cut off from civilization for months, Jack hopes to battle alcoholism and uncontrolled rage while writing a play. Evil forces residing in the Overlook – which has a long and violent history – covet young Danny for his precognitive powers and exploit Jack’s weaknesses to try to claim the boy.'),
       ('The Stand', 8, 1979, 'the-stand',
        E'One man escapes from a biological weapon facility after an accident, carrying with him the deadly virus known as Captain Tripps, a rapidly mutating flu that - in the ensuing weeks - wipes out most of the world\'s population. In the aftermath, survivors choose between following an elderly black woman to Boulder or the dark man, Randall Flagg, who has set up his command post in Las Vegas. The two factions prepare for a confrontation between the forces of good and evil.'),
       ('The Gunslinger', 8, 1982, 'the-gunslinger',
        E'The opening chapter in the epic Dark Tower series. Roland, the last gunslinger, in a world where time has moved on, pursues his nemesis, The Man in Black, across a desert. Roland\'s ultimate goal is the Dark Tower, the nexus of all universes. This mysterious icon\'s power is failing, threatening everything in existence.');

INSERT INTO mybooks (title, publisher_id, publication_year, slug, description)
VALUES ('IT', 8, 1986, 'it',
        E'A promise made twenty-eight years ago calls seven adults to reunite in Derry, Maine, where as teenagers they battled an evil creature that preyed on the city\'s children. Unsure that their Losers Club had vanquished the creature all those years ago, the seven had vowed to return to Derry if IT should ever reappear. Now, children are being murdered again and their repressed memories of that summer return as they prepare to do battle with the monster lurking in Derry\'s sewers once more.'),
       ('The Dead Zone', 8, 1979, 'the-dead-zone',
        E'Waking up from a five-year coma after a car accident, former schoolteacher Johnny Smith discovers that he can see people\'s futures and pasts when he touches them. Many consider his talent a gift; Johnny feels cursed. His fiance married another man during his coma and people clamor for him to solve their problems. When Johnny has a disturbing vision after he shakes the hand of an ambitious and amoral politician, he must decide if he should take drastic action to change the future.'),
       (E'\'Salem\'s Lot', 8, 1975, 'salems-lot',
        E'Author Ben Mears returns to ‘Salem\'s Lot to write a book about a house that has haunted him since childhood only to find his isolated hometown infested with vampires. While the vampires claim more victims, Mears convinces a small group of believers to combat the undead.');

INSERT INTO mybooks_authors (book_id, author_id)
VALUES (1, 1),
       (2, 1),
       (3, 1);

INSERT INTO mybooks_authors (book_id, author_id)
VALUES (4, 1),
       (5, 1),
       (6, 1);

INSERT INTO mybooks_genres (book_id, genre_id)
VALUES (1, 6),
       (1, 4),
       (2, 2),
       (2, 6),
       (3, 4);

INSERT INTO mybooks_genres (book_id, genre_id)
VALUES (4, 4),
       (5, 6),
       (6, 6);