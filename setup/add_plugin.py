#!/bin/python

if __name__ == '__main__':
    infile = open('plugin.go', 'r')
    infile_list = infile.readlines()
    start_index = 1e8 # large number
    end_index = -1
    for line_num, line in enumerate(infile_list):
        if line.startswith('var directives'):
            start_index = line_num
        elif line_num > start_index and line.strip() == '}':
            end_index = line_num
    if end_index != -1:
        new_list = infile_list[:end_index-1]
        new_list.append('\t"mux",\n')
        new_list.extend(infile_list[end_index:])
        infile.close()
        new_file = open('plugin.go', 'w')
        new_file.writelines(new_list)
    else:
        print 'Directives list not found'
